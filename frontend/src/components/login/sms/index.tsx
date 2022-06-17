import style from './index.less';
import { Form, Input, Select, Button } from 'antd';
import type { LoginProp } from '..';

const SMS: React.FC<LoginProp> = () => {
  return (
    <Form name="sms" className={style.form} size="large">
      <Form.Item className={style.telLine}>
        <Input.Group compact>
          <Form.Item name="area">
            <Select className={style.area} defaultValue="china">
              <Select.Option value="china">+1111</Select.Option>
              <Select.Option value="america">+1</Select.Option>
            </Select>
          </Form.Item>
          <Form.Item name="tel">
            <Input className={style.tel} placeholder="请输入手机号" />
          </Form.Item>
          <Button type="link">获取验证码</Button>
        </Input.Group>
      </Form.Item>
      <Form.Item className={style.codeLine}>
        <Input.Group compact>
          <span className={style.codeLabel}>验证码</span>
          <Form.Item name="code">
            <Input className={style.code} placeholder="请输入验证码" />
          </Form.Item>
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
export default SMS;
