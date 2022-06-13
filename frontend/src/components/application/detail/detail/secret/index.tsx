import style from './index.less';
import copy from 'copy-to-clipboard';
import { Button, message } from 'antd';
import { Secret } from '@/api/application/v1/application';
import { KeyOutlined, CopyOutlined } from '@ant-design/icons';
import moment from 'moment';
interface SecretProp {
  secret: Secret
  only: boolean
  isNew: boolean
}
const SecretItem: React.FC<SecretProp> = (prop) => {
  const copyClientSecret = (): void => {
      copy(prop.secret.maskedSecret);
      message.success(`已复制到剪切板: ${prop.secret.maskedSecret}`)
  }
  return (
    <div className={style.secret+(prop.isNew?` ${style.new}`:'')}>
      <div className={style.iconArea}>
        <KeyOutlined className={style.icon} />
        <span className={style.label}>客户端密码</span>
      </div>
      <div className={style.infoArea}>
        <span className={style.description}>说明: {prop.secret.description}</span>
        <div>
          <span className={style.clientSecret}>{prop.secret.maskedSecret}</span>
          { prop.isNew && <Button type='link' shape='circle' icon={<CopyOutlined />} size='small' onClick={copyClientSecret}></Button> }
        </div>
        <span className={style.generated}>生成时间: {moment(prop.secret.generated).calendar()}</span>
        <span className={style.lastUsed}>最后使用: {prop.secret.lastUsed?moment(prop.secret.lastUsed).calendar():'从未使用'}</span>
        { prop.only &&<span className={style.only}>您不能删除唯一的客户端密码。首先生成一个新的客户端密码。</span> }
      </div>
      <div className={style.actionArea}>
        <Button className={style.delete} size='large' danger>删除</Button>
      </div>
    </div>
  )
}
export default SecretItem;
