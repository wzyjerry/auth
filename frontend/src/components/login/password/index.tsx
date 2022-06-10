import style from './index.less';
import { Form, Input, Button, message } from 'antd';
import { LoginClient } from '@/api';
import { LoginRequest, Type, Method } from '@/api/user/v1/login_pb';
import { LoginForm } from '..';
import { Go, Reject } from '@/util';
interface passwordForm {
  account: string,
  password: string
}
const Password: React.FC<LoginForm> = (props) => {
  const onLogin = async (form: passwordForm): Promise<void> => {
    const client = new LoginClient();
    const request = new LoginRequest();
    request.setMethod(Method.METHOD_PASSWORD);
    request.setSecret(form.password);
    if (form.account.includes('@')) {
      request.setType(Type.TYPE_EMAIL);
      request.setUnique(form.account);
    } else {
      request.setType(Type.TYPE_PHONE);
      request.setUnique(`+86${form.account}`)
    }
    const reply = await Go(client.Login(request))
    if (reply instanceof Reject) {
      switch (reply.error.name) {
      case 'PASSWORD_LOGIN':
        message.error('用户名或密码错误');
        break;
      default:
        message.error(`[${reply.error.name}]${reply.error.message}`);
      }
    } else {
      props.onSuccess(reply.val.token);
    }
  }
  return (
    <Form
      name='password'
      className={style.form}
      onFinish={onLogin}
      size='large'
    >
      <Form.Item className={style.accountLine}>
        <Input.Group compact>
          <span className={style.accountLabel}>账号</span>
          <Form.Item name='account' rules={[{ required: true, message: '请输入邮箱/手机号'}]}>
            <Input className={style.account} placeholder='请输入邮箱/手机号'></Input>
          </Form.Item>
        </Input.Group>
      </Form.Item>
      <Form.Item className={style.passwordLine}>
        <Input.Group compact>
          <span className={style.passwordLabel}>密码</span>
          <Form.Item name='password' rules={[{ required: true, message: '请输入密码'}]}>
            <Input.Password className={style.password} placeholder='请输入密码'></Input.Password>
          </Form.Item>
          <Button type="link">忘记密码?</Button>
        </Input.Group>
      </Form.Item>
      <Form.Item className={style.submitLine}>
        <Button className={style.register}>注册</Button>
        <Button className={style.login} htmlType='submit' type='primary'>登录</Button>
      </Form.Item>
    </Form>
  );
}
export default Password;
