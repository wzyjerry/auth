import style from './index.less';
import { useHistory, useLocation } from 'react-router-dom';
import { Form, Input, Button, message } from 'antd';
import { CreateRequest } from '@/api/application/v1/application_pb';
import { ApplicationClient } from '@/api';
import { Go, Reject } from '@/util';
interface registerForm {
  name: string
  homepage: string
  description?: string
  callback: string
}
const New: React.FC = () => {
  const history = useHistory();
  const location = useLocation();
  const onRegister = async (form: registerForm): Promise<void> => {
    const client = new ApplicationClient();
    const request = new CreateRequest();
    request.setName(form.name);
    request.setHomepage(form.homepage);
    if (form.description !== undefined) {
      request.setDescription(form.description);
    }
    request.setCallback(form.callback);
    const reply = await Go(client.Create(request))
    if (reply instanceof Reject) {
      switch (reply.error.name) {
        case 'FORBIDDEN':
        case 'UNAUTHORIZED':
          history.push(`/user/login?return_to=${encodeURIComponent(location.pathname+location.search+location.hash)}`)
          break
        default:
          console.log(reply.error);
          message.error(`[${reply.error.name}]${reply.error.message}`);
      }
      return
    }
    history.push(`/application/${reply.val.id}`);
  }
  const onCancelClick = (): void => {
    history.push('/application');
  }
  return (
    <div className={style.new}>
      <div className={style.title}>
        <span>注册新的OAuth应用程序</span>
      </div>
      <hr className={style.line}></hr>
      <Form className={style.form} onFinish={onRegister} name='new' size='large' layout='vertical'>
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
      </Form>
    </div>
  )
}

export default New;
