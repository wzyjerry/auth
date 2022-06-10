import { Reject, Go } from '@/util';
import { Avatar } from 'antd';
import { Avatar as AvatarHelper, Token, ErrKeyNotExists } from '@/util/localStorage';
import { ProfileClient } from '@/api';
import { useState } from 'react';

export function IsLogin(): boolean {
  const helper = new Token();
  const token = helper.Load();
  if (token instanceof Reject) {
    switch (token.error) {
      case ErrKeyNotExists:
        return false
      default:
        console.error(token.error)
        return false
    }
  }
  return true;
}

export async function FetchAvatar():Promise<string> {
  const helper = new AvatarHelper();
  if (!IsLogin()) {
    return '';
  }
  const avatar = helper.Load();
  if (avatar instanceof Reject) {
    switch (avatar.error) {
    case ErrKeyNotExists:
      const client = new ProfileClient()
      const reply = await Go(client.GetAvatar())
      if (reply instanceof Reject) {
        return '';
      }
      helper.Save(reply.val.avatar);
      return reply.val.avatar;
    default:
      console.error(avatar.error)
      return '';
    }
  }
  return avatar.val
}

const Navbar: React.FC = () => {
  const [ avatar, setAvatar ] = useState<string>();
  (async() => {
    const avatar = await FetchAvatar();
    console.log(avatar)
    setAvatar(avatar);
  })()
  return (
    <Avatar size='large' src={avatar}></Avatar>
  )
}

export default Navbar;
