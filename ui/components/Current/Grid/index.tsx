import { FC } from 'react';
import useSWR from 'swr';

interface ResponseCurrent {
  temperature: number;
  humidity: number;
  timestamp: string; // '2021-01-10T13:25:48Z'
}

export const GridContainer: FC = ({ children }) => {
  return (
    <>
      <div className="grid">{children}</div>
      <style jsx>
        {`
          .grid {
            display: flex;
            align-items: center;
            justify-content: space-between;
            flex-wrap: wrap;
            width: 100%;
            margin-bottom: 16px;
          }
        `}
      </style>
    </>
  );
};

export const GridItem: FC = ({ children }) => {
  return (
    <>
      <div className="grid-item">{children}</div>
      <style jsx>
        {`
          .grid-item {
            width: 48%;
          }

          @media (max-width: 767px) {
            .grid-item {
              min-width: 160px;
            }
          }
        `}
      </style>
    </>
  );
};
