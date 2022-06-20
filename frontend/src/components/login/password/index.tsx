import style from './index.less';
import { Form, Input, Button } from 'antd';
import type { LoginRequest } from '@/api/user/v1/login';
import { Type, Method } from '@/api/user/v1/login';
import type { LoginProp } from '..';
import { useDispatch } from 'umi';
interface passwordForm {
  account: string;
  password: string;
}
const Password: React.FC<LoginProp> = ({ returnTo }) => {
  const dispatch = useDispatch();
  const onLogin = async (form: passwordForm): Promise<void> => {
    const request: LoginRequest = {
      method: Method.METHOD_PASSWORD,
      secret: form.password,
      type: Type.TYPE_UNSET,
    };
    if (form.account.includes('@')) {
      request.type = Type.TYPE_EMAIL;
      request.unique = form.account;
    } else {
      request.type = Type.TYPE_PHONE;
      request.unique = `+86${form.account}`;
    }
    dispatch({
      type: 'user/login',
      payload: {
        request: request,
        returnTo: returnTo,
      },
    });
  };
  return (
    <Form className={style.form} onFinish={onLogin} size="large">
      <Form.Item className={style.accountLine}>
        <Input.Group compact>
          <span className={style.accountLabel}>账号</span>
          <Form.Item name="account" rules={[{ required: true, message: '请输入邮箱/手机号' }]}>
            <Input className={style.account} placeholder="请输入邮箱/手机号" />
          </Form.Item>
        </Input.Group>
      </Form.Item>
      <Form.Item className={style.passwordLine}>
        <Input.Group compact>
          <span className={style.passwordLabel}>密码</span>
          <Form.Item name="password" rules={[{ required: true, message: '请输入密码' }]}>
            <Input.Password className={style.password} placeholder="请输入密码" />
          </Form.Item>
          <Button type="link">忘记密码?</Button>
        </Input.Group>
      </Form.Item>
      <Form.Item className={style.submitLine}>
        <Button className={style.register}>注册</Button>
        <Button className={style.login} htmlType="submit" type="primary">
          登录
        </Button>
      </Form.Item>
    </Form>
  );
};
export default Password;
