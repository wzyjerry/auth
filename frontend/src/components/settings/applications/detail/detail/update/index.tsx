import style from './index.less';
import { Form, Input, Button } from 'antd';
import { useDispatch, useSelector } from 'umi';
import { UpdateRequest } from '@/api/application/v1/application';
import type { Application } from 'umi';
interface applicationForm {
  name: string;
  homepage: string;
  description?: string;
  callback: string;
}
const UpdateForm: React.FC = () => {
  const dispatch = useDispatch();
  const application: Application = useSelector(
    ({ application }: { application: Application }) => application,
  );
  const fields = Object.entries(UpdateRequest.fromJSON(application))
    .map((x) => ({
      name: x[0],
      value: x[1],
    }))
    .filter((x) => x.name != 'id');
  const update = async (form: applicationForm): Promise<void> => {
    const request: Omit<UpdateRequest, 'id'> = {
      name: form.name,
      homepage: form.homepage,
      description: form.description,
      callback: form.callback,
    };
    if (request.description === '') {
      request.description = undefined;
    }
    dispatch({
      type: 'application/update',
      payload: request,
    });
  };
  return (
    <Form className={style.form} size="large" layout="vertical" fields={fields} onFinish={update}>
      <Form.Item className={style.editArea}>
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
      </Form.Item>
      <hr className={style.line} />
      <Form.Item className={style.submitLine}>
        <Button className={style.update} htmlType="submit" type="primary">
          更新应用
        </Button>
      </Form.Item>
    </Form>
  );
};
export default UpdateForm;
