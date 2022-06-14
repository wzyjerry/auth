import style from './index.less';
import { Modal, Form, Input, Button } from 'antd';
import { GenerateClientSecretRequest, GenerateClientSecretReply } from '@/api/application/v1/application';
import { ApplicationClient } from '@/api';
interface GenerateClientSecretProp {
  id: string
  visiable: boolean
  onClose: (promise?: Promise<GenerateClientSecretReply>) => void
}
interface generateClientSecretForm {
  description: string
}
const GenerateClientSecret: React.FC<GenerateClientSecretProp> = (prop) => {
  const onGenerate = async (form: generateClientSecretForm): Promise<void> => {
    const client = new ApplicationClient();
    const request: GenerateClientSecretRequest = {
      id: prop.id,
      description: form.description
    };
    prop.onClose(client.GenerateClientSecret(request));
  }
  const onClose = () => {
    prop.onClose();
  }
  return (
    <Modal className={style.modal} destroyOnClose footer={null} visible={prop.visiable} onCancel={onClose}>
      <Form className={style.form} onFinish={onGenerate} name='new' size='large' layout='vertical'>
        <Form.Item label='说明' name='description' rules={[{ required: true, message: '请简要说明密码用途'}]}>
          <Input className={style.description}></Input>
        </Form.Item>
        <Form.Item className={style.submitLine}>
          <Button className={style.register} htmlType='submit' type='primary'>生成</Button>
          <Button className={style.cancel} onClick={onClose}>取消</Button>
        </Form.Item>
      </Form>
    </Modal>
  )
}
export default GenerateClientSecret;
