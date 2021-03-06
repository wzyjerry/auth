import style from './index.less';
import { Tabs } from 'antd';
import Password from './password';
import SMS from './sms';
import ThirdParty from './third_party';
export interface LoginProp {
  returnTo: string;
}
const Login: React.FC<LoginProp> = ({ returnTo }) => {
  return (
    <div className={style.login}>
      <Tabs className={style.tab} centered size="large">
        <Tabs.TabPane tab="密码登录" key="1">
          <Password returnTo={returnTo} />
        </Tabs.TabPane>
        <Tabs.TabPane tab="短信登录" key="2">
          <SMS returnTo={returnTo} />
        </Tabs.TabPane>
      </Tabs>
      <ThirdParty className={style.thirdParty} returnTo={returnTo} />
    </div>
  );
};
export default Login;
