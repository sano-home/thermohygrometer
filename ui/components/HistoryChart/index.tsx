import { FC } from 'react';
import useSWR from 'swr';
import { ResponsiveLine } from '@nivo/line';

import { ResponseCurrent } from '../Current';
import { colors } from '../../constants/colors';

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
  return histories.map((item, index) => ({
    x: formatTimestamp(item.timestamp),
    y: item[field],
  }));
};

export const HistoryChart: FC = () => {
  const { data, error } = useSWR<ResponseHistories, Error>('/api/histories');

  if (error) return <div>failed to load</div>;
  if (!data) return <div>loading...</div>;

  const temperatureChartData: LineChartItem = {
    id: 'temperature',
    data: convertToChartData(data.data, 'temperature'),
  };

  const humidityChartData: LineChartItem = {
    id: 'humidity',
    data: convertToChartData(data.data, 'humidity'),
  };

  return (
    <div style={{ height: 320, width: 800 }}>
      <ResponsiveLine
        colors={[colors.temperature]}
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
        colors={[colors.humidity]}
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
    </div>
  );
};

const formatTimestamp = (timestamp: string): string => {
  const time = new Date(timestamp);
  const hour = time.getHours().toString().padStart(2, '0');
  const minute = time.getMinutes().toString().padStart(2, '0');
  return `${hour}:${minute}`;
};
