import { FC } from 'react';
import useSWR from 'swr';
import { ResponsiveLine } from '@nivo/line';

import { ResponseCurrent } from '../Current';

interface ResponseHistories {
  pages: {
    total: number;
  };
  data: ResponseCurrent[];
}

interface LineChartItem {
  id: string;
  color: string;
  data: {
    x: string;
    y: number;
  }[];
}

const convertToChartData = (
  histories: ResponseHistories['data'],
  field: 'temperature' | 'humidity'
): LineChartItem['data'] => {
  return histories.map((item, index) => ({
    x: item.timestamp,
    y: item[field],
  }));
};

export const HistoryChart: FC = () => {
  const { data, error } = useSWR<ResponseHistories, Error>('/api/histories');

  if (error) return <div>failed to load</div>;
  if (!data) return <div>loading...</div>;

  const temperatureChartData: LineChartItem = {
    id: 'temperature',
    color: 'hsl(190, 70%, 50%)',
    data: convertToChartData(data.data, 'temperature'),
  };

  const humidityChartData: LineChartItem = {
    id: 'humidity',
    color: 'hsl(207, 70%, 50%)',
    data: convertToChartData(data.data, 'humidity'),
  };

  return (
    <div style={{ height: 320, width: 800 }}>
      <ResponsiveLine
        data={[temperatureChartData]}
        margin={{ top: 50, right: 40, bottom: 50, left: 80 }}
        xScale={{ type: 'point' }}
        yScale={{
          type: 'linear',
          min: 'auto',
          max: 'auto',
          stacked: true,
          reverse: false,
        }}
        yFormat=" >-.2f"
        axisTop={null}
        axisRight={null}
        axisBottom={{
          orient: 'bottom',
          tickSize: 5,
          tickPadding: 5,
          tickRotation: 0,
          legend: 'time',
          legendOffset: 36,
          legendPosition: 'middle',
        }}
        axisLeft={{
          orient: 'left',
          tickSize: 5,
          tickPadding: 5,
          tickRotation: 0,
          legend: 'temperature',
          legendOffset: -60,
          legendPosition: 'middle',
        }}
        pointSize={10}
        pointColor={{ theme: 'background' }}
        pointBorderWidth={2}
        pointBorderColor={{ from: 'serieColor' }}
        pointLabelYOffset={-12}
        useMesh={true}
      />

      <ResponsiveLine
        data={[humidityChartData]}
        margin={{ top: 50, right: 40, bottom: 50, left: 80 }}
        xScale={{ type: 'point' }}
        yScale={{
          type: 'linear',
          min: 'auto',
          max: 'auto',
          stacked: true,
          reverse: false,
        }}
        yFormat=" >-.2f"
        axisTop={null}
        axisRight={null}
        axisBottom={{
          orient: 'bottom',
          tickSize: 5,
          tickPadding: 5,
          tickRotation: 0,
          legend: 'time',
          legendOffset: 36,
          legendPosition: 'middle',
        }}
        axisLeft={{
          orient: 'left',
          tickSize: 5,
          tickPadding: 5,
          tickRotation: 0,
          legend: 'humidity',
          legendOffset: -60,
          legendPosition: 'middle',
        }}
        pointSize={10}
        pointColor={{ theme: 'background' }}
        pointBorderWidth={2}
        pointBorderColor={{ from: 'serieColor' }}
        pointLabelYOffset={-12}
        useMesh={true}
      />

      {data.data.map((item, index) => (
        <div key={item.timestamp}>
          <p>{item.timestamp}</p>
          <p>{item.temperature}</p>
          <p>{item.humidity}</p>
          <hr />
        </div>
      ))}
    </div>
  );
};
