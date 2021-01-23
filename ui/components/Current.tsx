import { FC } from 'react';
import useSWR from 'swr';

interface ResponseCurrent {
  temperature: number;
  humidity: number;
  timestamp: string; // '2021-01-10T13:25:48Z'
}

export const Current: FC = () => {
  const { data, error } = useSWR<ResponseCurrent, Error>('/api/current');
  console.log(data, error);

  if (error) return <div>failed to load</div>;
  if (!data) return <div>loading...</div>;

  // render data
  return (
    <div>
      <p>気温：{data.temperature}</p>
      <p>湿度：{data.humidity}</p>
    </div>
  );
};
