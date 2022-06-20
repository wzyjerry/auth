import { ApplicationDetail } from '@/components';
import { useDispatch, useParams } from 'umi';
import { useEffect } from 'react';
const Detail: React.FC = () => {
  const dispatch = useDispatch();
  const { id } = useParams<{ id: string }>();
  useEffect(() => {
    dispatch({
      type: 'application/setup',
      payload: { id },
    });
  }, [dispatch, id]);
  return <ApplicationDetail />;
};
export default Detail;
