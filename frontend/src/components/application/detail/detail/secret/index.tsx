import style from './index.less';
import copy from 'copy-to-clipboard';
import { Button, Popconfirm, message } from 'antd';
import type { Secret } from '@/api/application/v1/application';
import { KeyOutlined, CopyOutlined } from '@ant-design/icons';
import moment from 'moment';
import { ObjectID } from 'bson';
import { useDispatch } from 'umi';
interface SecretProp {
  secret: Secret;
  only: boolean;
}
const SecretItem: React.FC<SecretProp> = (prop) => {
  const dispatch = useDispatch();
  const copyClientSecret = (): void => {
    copy(prop.secret.secret);
    message.success(`已复制到剪切板: ${prop.secret.secret}`);
  };
  const onRevoke = () => {
    dispatch({
      type: 'application/revokeClientSecret',
      payload: {
        secretId: prop.secret.id,
      },
    });
  };
  return (
    <div className={style.secret + (!prop.secret.masked ? ` ${style.new}` : '')}>
      <div className={style.iconArea}>
        <KeyOutlined className={style.icon} />
        <span className={style.label}>客户端密钥</span>
      </div>
      <div className={style.infoArea}>
        <span className={style.description}>说明: {prop.secret.description}</span>
        <div>
          <span className={style.clientSecret}>{prop.secret.secret}</span>
          {prop.secret.masked || (
            <Button
              type="link"
              shape="circle"
              icon={<CopyOutlined />}
              size="small"
              onClick={copyClientSecret}
            />
          )}
        </div>
        <span className={style.generated}>
          生成于 {moment(new ObjectID(prop.secret.id).getTimestamp()).calendar()}
        </span>
        <span className={style.lastUsed}>
          最后使用时间:{' '}
          {prop.secret.lastUsed ? moment(prop.secret.lastUsed).calendar() : '从未使用'}
        </span>
        {prop.only && (
          <span className={style.only}>
            您不能删除唯一的客户端密钥。首先生成一个新的客户端密钥。
          </span>
        )}
      </div>
      <div className={style.actionArea}>
        <Popconfirm
          placement="topRight"
          title="此操作无法撤消。此客户端密钥将立即停止工作。您确定要删除此客户端密钥吗？"
          onConfirm={onRevoke}
          okText="确定"
          cancelText="取消"
          disabled={prop.only}
        >
          <Button className={style.revoke} size="large" danger disabled={prop.only}>
            删除
          </Button>
        </Popconfirm>
      </div>
    </div>
  );
};
export default SecretItem;
