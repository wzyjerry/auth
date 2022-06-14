import style from './index.less';
import copy from 'copy-to-clipboard';
import { Button, Popconfirm, message } from 'antd';
import { Secret, RevokeClientSecretRequest } from '@/api/application/v1/application';
import { KeyOutlined, CopyOutlined } from '@ant-design/icons';
import moment from 'moment';
import { ObjectID } from 'bson';
import { ApplicationClient } from '@/api';
import { Go, Reject } from '@/util';
import { useHistory } from 'react-router-dom';
interface SecretProp {
  id: string
  secret: Secret
  only: boolean
  isNew: boolean
  refresh: ()=>void
}
const SecretItem: React.FC<SecretProp> = (prop) => {
  const history = useHistory();
  const copyClientSecret = (): void => {
      copy(prop.secret.maskedSecret);
      message.success(`已复制到剪切板: ${prop.secret.maskedSecret}`)
  }
  const onRevoke = async (): Promise<void> => {
    const client = new ApplicationClient();
    const request: RevokeClientSecretRequest = {
      id: prop.id,
      secretId: prop.secret.id,
    }
    const reply = await Go(client.RevokeClientSecret(request))
    if (reply instanceof Reject) {
      switch(reply.error.name) {
        case 'FORBIDDEN':
          case 'UNAUTHORIZED':
            history.push(`/user/login?return_to=${encodeURIComponent(location.pathname+location.search+location.hash)}`);
            break;
          case 'APPLICATION_NOT_FOUND':
            history.push(`/404`);
            break;
          case 'REVOKE_BAD_REQUEST':
            message.error(`不能删除最后一个客户端密钥`);
            break;
          default:
            console.log(reply.error);
            message.error(`[${reply.error.name}]${reply.error.message}`);
      }
    }
    message.success('客户端密钥已删除');
    prop.refresh()
  }
  return (
    <div className={style.secret+(prop.isNew?` ${style.new}`:'')}>
      <div className={style.iconArea}>
        <KeyOutlined className={style.icon} />
        <span className={style.label}>客户端密钥</span>
      </div>
      <div className={style.infoArea}>
        <span className={style.description}>说明: {prop.secret.description}</span>
        <div>
          <span className={style.clientSecret}>{prop.secret.maskedSecret}</span>
          { prop.isNew && <Button type='link' shape='circle' icon={<CopyOutlined />} size='small' onClick={copyClientSecret}></Button> }
        </div>
        <span className={style.generated}>生成于 {moment(new ObjectID(prop.secret.id).getTimestamp()).calendar()}</span>
        <span className={style.lastUsed}>最后使用时间: {prop.secret.lastUsed?moment(prop.secret.lastUsed).calendar():'从未使用'}</span>
        { prop.only &&<span className={style.only}>您不能删除唯一的客户端密钥。首先生成一个新的客户端密钥。</span> }
      </div>
      <div className={style.actionArea}>
        <Popconfirm
          placement='topRight'
          title='此操作无法撤消。此客户端密钥将立即停止工作。您确定要删除此客户端密钥吗？'
          onConfirm={onRevoke}
          okText="确定"
          cancelText="取消"
          disabled={prop.only}
        >
          <Button className={style.revoke} size='large' danger disabled={prop.only}>删除</Button>
        </Popconfirm>
      </div>
    </div>
  )
}
export default SecretItem;
