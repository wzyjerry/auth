import style from './index.less';
import copy from 'copy-to-clipboard';
import { CopyOutlined } from '@ant-design/icons';
import { Button, message, Alert } from 'antd';
import { useHistory } from 'react-router-dom';
import { RetrieveRequest, RetrieveReply, GenerateClientSecretReply, Secret } from '@/api/application/v1/application';
import { ApplicationClient } from '@/api';
import { Go, Reject } from '@/util';
import { useEffect, useState } from 'react';
import SecretItem from './secret';
import GenerateClientSecret from './generateClientSecret';
interface DetailProp {
  id: string
}
const Detail: React.FC<DetailProp> = (prop) => {
  const history = useHistory();
  let [ detail, setDetail ] = useState<RetrieveReply>();
  let [ refresh, setRefresh ] = useState(false);
  const refreshDetail = (): void => {
    setRefresh(!refresh);
  }
  useEffect(() => {
    (async () => {
      const client = new ApplicationClient();
      const request: RetrieveRequest = {
        id: prop.id
      };
      const reply = await Go(client.Retrieve(request))
      if (reply instanceof Reject) {
        switch (reply.error.name) {
          case 'FORBIDDEN':
          case 'UNAUTHORIZED':
            history.push(`/user/login?return_to=${encodeURIComponent(location.pathname+location.search+location.hash)}`)
            break
          case 'APPLICATION_NOT_FOUND':
            history.push(`/404`)
            break
          default:
            console.log(reply.error);
            message.error(`[${reply.error.name}]${reply.error.message}`);
        }
        return
      }
      setDetail(reply.val);
    })()
  }, [prop.id, refresh])
  const copyClientId = (): void => {
    if (detail) {
      copy(detail.clientId);
      message.success(`已复制到剪切板: ${detail.clientId}`)
    }
  }
  const isNew = (secret: Secret): boolean => {
    return !secret.maskedSecret.startsWith('*');
  }
  const renderClientSecrets = (): JSX.Element => {
    if (!detail) {
      return (<></>)
    }
    if (detail.clientSecrets.length > 0) {
      const only = detail.clientSecrets.length === 1;
      return (
        <div className={style.clientSecretArea}>
        {
          isNew(detail.clientSecrets[0]) && <Alert className={style.alert} type='info' description='确保现在复制您的新客户端密钥。你将无法再次看到它。'></Alert>
        }
          <div className={style.clientSecretList}>
            {
              detail.clientSecrets.map((secret) => <SecretItem id={prop.id} secret={secret} only={only} isNew={isNew(secret)} key={secret.id} refresh={refreshDetail}></SecretItem>)
            }
          </div>
        </div>

      )
    } else {
      return <span>您需要一个客户端密钥来验证应用程序。</span>
    }
  }
  const [visible, setVisible] = useState(false);
  const generateClientSecret = async (promise?: Promise<GenerateClientSecretReply>): Promise<void> => {
    if (promise !== undefined) {
      const reply = await Go(promise);
      if (reply instanceof Reject) {
        switch (reply.error.name) {
          case 'FORBIDDEN':
          case 'UNAUTHORIZED':
            history.push(`/user/login?return_to=${encodeURIComponent(location.pathname+location.search+location.hash)}`)
            break
          case 'APPLICATION_NOT_FOUND':
            history.push(`/404`)
            break
          default:
            console.log(reply.error);
            message.error(`[${reply.error.name}]${reply.error.message}`);
        }
        return
      }
      if (reply.val.secret && detail) {
        if (detail.clientSecrets) {
          detail.clientSecrets = Array<Secret>(reply.val.secret, ...detail.clientSecrets);
        } else {
          detail.clientSecrets = Array<Secret>(reply.val.secret);
        }
        setDetail(detail);
      }
    }
    setVisible(false);
  }
  return (
    <div className={style.detail}>
      <div className={style.title}>
        <span>{detail?.name}</span>
      </div>
      <hr className={style.line}></hr>
      <div className={style.clientIdLine}>
        <span className={style.clientIdLabel}>客户端ID:</span>
        <span className={style.clientId}>{detail?.clientId}</span>
        <Button type='link' shape='circle' icon={<CopyOutlined />} size='large' onClick={copyClientId}></Button>
      </div>
      <div className={style.clientSecretLine}>
        <span className={style.clientSecretLabel}>客户端密码</span>
        <Button className={style.generate} size='large' onClick={()=>{setVisible(true)}}>生成新的客户端密码</Button>
        <GenerateClientSecret id={prop.id} visiable={visible} onClose={generateClientSecret}></GenerateClientSecret>
      </div>
      {renderClientSecrets()}
      {/* <Form className={style.form} onFinish={onRegister} name='new' size='large' layout='vertical'>
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
      </Form> */}
    </div>
  )
}
export default Detail;
