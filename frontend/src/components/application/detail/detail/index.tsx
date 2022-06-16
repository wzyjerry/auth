import style from './index.less';
import copy from 'copy-to-clipboard';
import { CopyOutlined } from '@ant-design/icons';
import { Button, message, Alert } from 'antd';
import { useState } from 'react';
import SecretItem from './secret';
import GenerateClientSecret from './generateClientSecret';
import { useSelector, Application } from 'umi';

const Detail: React.FC = () => {
  const application: Application = useSelector(state=>state.application);
  const copyClientId = (): void => {
    copy(application.clientId);
    message.success(`已复制到剪切板: ${application.clientId}`)
  }
  const renderClientSecrets = (): JSX.Element => {
    if (application.clientSecrets.length > 0) {
      const only = application.clientSecrets.length === 1;
      return (
        <div className={style.clientSecretArea}>
        {
          application.clientSecrets[0].masked || <Alert className={style.alert} type='info' description='确保现在复制您的新客户端密钥。你将无法再次看到它。'></Alert>
        }
          <div className={style.clientSecretList}>
            {
              application.clientSecrets.map((secret) => <SecretItem key={secret.id} secret={secret} only={only}></SecretItem>)
            }
          </div>
        </div>
      )
    } else {
      return <span>您需要一个客户端密钥来验证应用程序。</span>
    }
  }
  const [visible, setVisible] = useState(false);
  return (
    <div className={style.detail}>
      <div className={style.title}>
        <span>{application.name}</span>
      </div>
      <hr className={style.line}></hr>
      <div className={style.clientIdLine}>
        <span className={style.clientIdLabel}>客户端ID:</span>
        <span className={style.clientId}>{application.clientId}</span>
        <Button type='link' shape='circle' icon={<CopyOutlined />} size='large' onClick={copyClientId}></Button>
      </div>
      <div className={style.clientSecretLine}>
        <span className={style.clientSecretLabel}>客户端密码</span>
        <Button className={style.generate} size='large' onClick={()=>setVisible(true)}>生成新的客户端密码</Button>
        <GenerateClientSecret id={application.id} visiable={visible} onClose={()=>setVisible(false)}></GenerateClientSecret>
      </div>
      {renderClientSecrets()}
      {/* <Form className={style.form} onFinish={onRegister} name='new' size='large' layout='vertical'>
        <Form.Item label='应用名称' name='name' rules={[{ required: true, message: '请输入可以被用户识别和信任的应用名称'}]}>
          <Input></Input>
        </Form.Item>
        <Form.Item label='应用主页' name='homepage' rules={[{ required: true, type: 'url', message: '请输入应用的完整主页'}]}>
          <Input></Input>
        </Form.Item>
        <Form.Item label='应用简介' name='description'>
          <Input.TextArea placeholder='应用简介是可选的'></Input.TextArea>
        </Form.Item>
        <Form.Item label='授权回调地址' name='callback' rules={[{ required: true, type: 'url', message: '请输入应用的回调地址'}]}>
          <Input></Input>
        </Form.Item>
        <hr className={style.line}></hr>
        <Form.Item className={style.submitLine}>
          <Button className={style.register} htmlType='submit' type='primary'>注册应用</Button>
          <Button className={style.cancel} onClick={onCancelClick}>取消</Button>
        </Form.Item>
      </Form> */}
    </div>
  )
}

export default Detail;
