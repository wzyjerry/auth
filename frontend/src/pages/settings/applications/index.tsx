import { useDispatch } from 'umi';
import { useEffect } from 'react';
import { ApplicationOverview } from '@/components';
const Index: React.FC = () => {
  const dispatch = useDispatch();
  useEffect(() => {
    dispatch({
      type: 'application/setupAll',
    });
  }, [dispatch]);
  return <ApplicationOverview />;
};
export default Index;
