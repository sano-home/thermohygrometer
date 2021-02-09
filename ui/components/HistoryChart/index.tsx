import { FC } from 'react';
import useSWR from 'swr';
import { ResponsiveLine } from '@nivo/line';

import { ResponseCurrent } from '../Current';
import { colors } from '../../constants/colors';
import { ChartContainer } from './ChartContainer';

interface ResponseHistories {
  pages: {
    total: number;
  };
  data: ResponseCurrent[];
}

interface LineChartItem {
  id: string;
  data: {
    x: string;
    y: number;
  }[];
}

const convertToChartData = (
  histories: ResponseHistories['data'],
  field: 'temperature' | 'humidity'
): LineChartItem['data'] => {
  return histories
    .slice(0, histories.length)
    .reverse()
    .map((item, index) => ({
      x: formatTimestamp(item.timestamp),
      y: item[field],
    }));
};

export const HistoryChart: FC = () => {
  const fetcher = (url: string) => {
    const now = new Date();
    const before = now.toISOString();
    const count = 12;
    const interval = 60 * 60 * 1000; // 60 minutes

    return fetch(
      `${url}?before=${before}&count=${count}&interval=${interval}`
    ).then((r) => r.json());
  };

  const { data, error } = useSWR<ResponseHistories, Error>(
    '/api/histories',
    fetcher
  );

  if (error) return <div>failed to load</div>;
  if (!data) return <div>loading histories...</div>;

  const temperatureChartData: LineChartItem = {
    id: 'temperature',
    data: convertToChartData(data.data, 'temperature'),
  };

  const humidityChartData: LineChartItem = {
    id: 'humidity',
    data: convertToChartData(data.data, 'humidity'),
  };

  return (
    <>
      <ChartContainer>
        <ResponsiveLine
          colors={[colors.temperature]}
          data={[temperatureChartData]}
          margin={{ top: 50, right: 40, bottom: 50, left: 80 }}
          xScale={{ type: 'point' }}
          yScale={{
            type: 'linear',
            min: 0,
            max: 30,
            stacked: true,
            reverse: false,
          }}
          axisTop={null}
          axisRight={{
            orient: 'left',
            tickSize: 5,
            tickPadding: 5,
            tickRotation: 0,
          }}
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
      </ChartContainer>
      <ChartContainer>
        <ResponsiveLine
          data={[humidityChartData]}
          colors={[colors.humidity]}
          margin={{ top: 50, right: 40, bottom: 50, left: 80 }}
          xScale={{ type: 'point' }}
          yScale={{
            type: 'linear',
            min: 0,
            max: 100,
            stacked: true,
            reverse: false,
          }}
          yFormat=" >-.0r"
          axisTop={null}
          axisRight={{
            orient: 'left',
            tickSize: 5,
            tickPadding: 5,
            tickRotation: 0,
          }}
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
      </ChartContainer>
    </>
  );
};

const formatTimestamp = (timestamp: string): string => {
  const time = new Date(timestamp);
  const hour = time.getHours().toString().padStart(2, '0');
  const minute = time.getMinutes().toString().padStart(2, '0');
  const second = time.getSeconds().toString().padStart(2, '0');
  return `${hour}:${minute}`;
};
