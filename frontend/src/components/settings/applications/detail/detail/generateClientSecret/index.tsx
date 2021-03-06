import style from './index.less';
import { Modal, Form, Input, Button } from 'antd';
import { useDispatch } from 'umi';
interface GenerateClientSecretProp {
  visiable: boolean;
  onClose: () => void;
}
interface generateClientSecretForm {
  description: string;
}
const GenerateClientSecret: React.FC<GenerateClientSecretProp> = ({ visiable, onClose }) => {
  const dispatch = useDispatch();
  const onGenerate = (form: generateClientSecretForm) => {
    dispatch({
      type: 'application/generateClientSecret',
      payload: {
        description: form.description,
      },
    });
    onClose();
  };
  return (
    <Modal
      className={style.modal}
      destroyOnClose
      footer={null}
      visible={visiable}
      onCancel={onClose}
    >
      <Form className={style.form} onFinish={onGenerate} name="new" size="large" layout="vertical">
        <Form.Item
          label="说明"
          name="description"
          rules={[{ required: true, message: '请简要说明密码用途' }]}
        >
          <Input className={style.description} />
        </Form.Item>
        <Form.Item className={style.submitLine}>
          <Button className={style.register} htmlType="submit" type="primary">
            生成
          </Button>
          <Button className={style.cancel} onClick={onClose}>
            取消
          </Button>
        </Form.Item>
      </Form>
    </Modal>
  );
};
export default GenerateClientSecret;
