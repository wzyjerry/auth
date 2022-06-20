import style from './index.less';
import { history, useDispatch } from 'umi';
import { Form, Input, Button } from 'antd';
import type { CreateRequest } from '@/api/application/v1/application';

interface applicationForm {
  name: string;
  homepage: string;
  description?: string;
  callback: string;
}
const New: React.FC = () => {
  const dispatch = useDispatch();
  const create = async (form: applicationForm): Promise<void> => {
    const request: CreateRequest = {
      name: form.name,
      homepage: form.homepage,
      description: form.description,
      callback: form.callback,
    };
    if (request.description === '') {
      request.description = undefined;
    }
    dispatch({
      type: 'application/create',
      payload: request,
    });
  };
  const onCancelClick = (): void => {
    history.push('/settings/applications');
  };
  return (
    <div className={style.new}>
      <div className={style.title}>
        <span>注册新的OAuth应用程序</span>
      </div>
      <hr className={style.line} />
      <Form className={style.form} onFinish={create} size="large" layout="vertical">
        <Form.Item
          label="应用名称"
          name="name"
          rules={[{ required: true, message: '请输入可以被用户识别和信任的应用名称' }]}
        >
          <Input />
        </Form.Item>
        <Form.Item
          label="应用主页"
          name="homepage"
          rules={[{ required: true, type: 'url', message: '请输入应用的完整主页' }]}
        >
          <Input />
        </Form.Item>
        <Form.Item label="应用描述" name="description">
          <Input.TextArea placeholder="应用描述是可选的" />
        </Form.Item>
        <Form.Item
          label="授权回调地址"
          name="callback"
          rules={[{ required: true, type: 'url', message: '请输入应用的回调地址' }]}
        >
          <Input />
        </Form.Item>
        <hr className={style.line} />
        <Form.Item className={style.submitLine}>
          <Button className={style.create} htmlType="submit" type="primary">
            注册应用
          </Button>
          <Button className={style.cancel} onClick={onCancelClick}>
            取消
          </Button>
        </Form.Item>
      </Form>
    </div>
  );
};

export default New;
