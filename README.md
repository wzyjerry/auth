# auth
身份认证微服务

## 综述
相比于第一版单后端的融合方案，该鉴权微服务添加了前端项目auth，以支撑OIDC和OAuth2抽象。

要使用该服务，用户首先需要注册应用程序：使用该微服务提供的任意一种方式完成登录，接下来注册应用程序并在管理页面生成至少一个客户端密钥；接下来便可以接入该服务认证用户身份。

身份认证微服务对外暴露了OIDC抽象和部分OAuth 2.0抽象，去除了对Password方式和Implicit方式的支持。第三方服务发起请求时，身份认证微服务首先引导用户登录（如果已登录则继续）；接下来尝试获得用户授权：报告应用程序详情以及请求的作用域；如果用户同意，则返回id_token和/或code、token，进行接下来的操作。

- id_token基于OIDC抽象，是一个JWT，具备用户基本信息。在token过期前，应用程序可以直接使用令牌所包含的用户信息而不需要额外的接口调用。
- code基于OAuth2抽象，是预登录签发的短时间一次性口令，应用程序可以使用它换取access_token。
- access_token基于OAuth2抽象，尽管也可以使用JWT格式，但是用户不应假定任何信息。只用于作为身份令牌调用作用域内的身份认证微服务接口。

## 登录类型和方式组合表

|Type\Method|UNSET|PASSWORD|CODE|
|---:|:---:|:---:|:---:|
|**UNSET**|-|-|-|
|**ACCOUNT**|-|用户密码登录|-|
|**EMAIL**|-|邮箱密码登录|邮箱验证码快速登录|
|**PHONE**|-|手机密码登录|短信验证码快速登录|
|**GITHUB**|-|-|Github OAuth2三方登录|
|**MICROSOFT**|-|-|Azure AD OAuth2三方登录|

- 使用Password登录：unique传对应类型的唯一值，secret传密码
- 使用Code登录：unique传对应类型的唯一值，secret传验证码；对于OAuth2三方登录，只需要secret传code

## OAuth2终结点和参数组合表
|scope|含义|
|---:|:---|
|user|允许调用个人信息接口|
|offline_access|签发refresh_token|
|openid|签发id_token|

- *注: 使用空格组合多个scope

由于不提供Implicit方式，因此实际上使用的response_type如下表所示

|方式|response_type|签发方式|
|:---|---:|:---|
|1|code|Authorize 签发code<br/>Token签发access_token|
|2|id_token|Authorize签发id_token|
|3|code id_token|Authorize 签发code和id_token<br/>Token签发access_token和id_token|

- *: 对于OAuth2登录，只允许response_type=code，即方式1
- **: 对于OIDC登录，response_type中必须包含id_token，scope中必须包含openid且必须包含nonce参数，即方式2、3

|方式|grant_type|含义|
|:--|---:|:---|
|1|authorization_code|授权码登录|
|2|client_credentials|客户端登录|
|3|refresh_token|刷新令牌|
