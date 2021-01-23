import { FC } from 'react';
import useSWR from 'swr';

import { Card } from './Card';
import { GridContainer, GridItem } from './Grid';

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
    <>
      <GridContainer>
        <GridItem>
          <Card
            title="Temperature"
            value={data.temperature.toString()}
            suffix="℃"
          />
        </GridItem>
        <GridItem>
          <Card title="Humidity" value={data.humidity.toString()} suffix="%" />
        </GridItem>
      </GridContainer>
    </>
  );
};