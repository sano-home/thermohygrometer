import { FC } from 'react';
import useSWR from 'swr';

interface ResponseCurrent {
  temperature: number;
  humidity: number;
  timestamp: string; // '2021-01-10T13:25:48Z'
}

export const Card: FC<{
  title: string;
  value: string;
  suffix: string;
}> = ({ title, value, suffix }) => {
  return (
    <>
      <div className="card">
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
            border: 1px solid #eaeaea;
            border-radius: 10px;
            font-weight: bold;
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
