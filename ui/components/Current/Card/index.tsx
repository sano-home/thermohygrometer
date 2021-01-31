import { FC } from 'react';
import useSWR from 'swr';

import { colors } from '../../../constants/colors';

interface ResponseCurrent {
  temperature: number;
  humidity: number;
  timestamp: string; // '2021-01-10T13:25:48Z'
}

export const Card: FC<{
  title: string;
  value: string;
  suffix: string;
  colorTheme: 'temperature' | 'humidity';
}> = ({ title, value, suffix, colorTheme }) => {
  return (
    <>
      <div className={`card ${colorTheme}`}>
        <h3>{title}</h3>
        <p>
          <span className="value">{value}</span>
          <span className="suffix">{suffix}</span>
        </p>
      </div>

      <style jsx>
        {`
          .card {
            width: 100%;
            flex-basis: 45%;
            padding: 1.5rem;
            text-align: left;
            color: inherit;
            border: 2px solid;
            border-radius: 10px;
            background-color: #fff;
            font-weight: bold;
          }

          .temperature {
            border-color: ${colors.temperature};
          }

          .humidity {
            border-color: ${colors.humidity};
          }

          h3 {
            margin: 0.5rem;
          }

          p {
            text-align: right;
            margin: 0;
          }

          .value {
            font-size: 4rem;
          }

          .suffix {
            font-size: 2rem;
          }
        `}
      </style>
    </>
  );
};
