import style from './index.less';
import { Form, Input, Select, Button } from 'antd';
const { Option } = Select;
import { LoginForm } from '..';

const SMS: React.FC<LoginForm> = (props) => {
  return (
    <Form
      name='sms'
      className={style.form}
    >
      <Form.Item className={style.telLine}>
        <Input.Group compact size='large'>
          <Form.Item name='area'>
            <Select className={style.area} size='large' defaultValue="china">
              <Option value="china">+1111</Option>
              <Option value="america">+1</Option>
            </Select>
          </Form.Item>
          <Form.Item name='tel'>
            <Input className={style.tel} placeholder='请输入手机号'></Input>
          </Form.Item>
          <Button size='large' type="link">获取验证码</Button>
        </Input.Group>
      </Form.Item>
      <Form.Item className={style.codeLine}>
        <Input.Group compact size='large'>
          <span className={style.codeLabel}>验证码</span>
          <Form.Item name='code'>
            <Input className={style.code} placeholder='请输入验证码'></Input>
          </Form.Item>
        </Input.Group>
      </Form.Item>
      <Form.Item className={style.submitLine}>
        <Button className={style.register} size='large'>注册</Button>
        <Button className={style.login} htmlType='submit' size='large' type='primary'>登录</Button>
      </Form.Item>
    </Form>
  );
}
export default SMS;
