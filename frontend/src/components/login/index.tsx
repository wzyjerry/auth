import styles from './index.less';
import { Tabs } from 'antd';
const {TabPane} = Tabs;
import Password from './password';
import SMS from './sms';
import ThirdParty from './third_party';
interface LoginFormProp {
  onSuccess: (token: string) => void
}
interface LoginProp extends LoginFormProp {
  returnTo: string
}
const Login: React.FC<LoginProp> = (prop) => {
  return (
    <div className={styles.login}>
      <Tabs className={styles.tab} centered size='large'>
        <TabPane tab="密码登录" key="1">
          <Password onSuccess={prop.onSuccess}></Password>
        </TabPane>
        <TabPane tab="短信登录" key="2">
          <SMS onSuccess={prop.onSuccess}></SMS>
        </TabPane>
      </Tabs>
      <ThirdParty className={styles.thirdParty} returnTo={prop.returnTo}></ThirdParty>
    </div>
  )
}

export { LoginFormProp as LoginForm };
export default Login;
