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
            display: grid;
            width: 100%;
            grid-template-columns: repeat(2, 1fr);
            grid-column-gap: 16px;
            grid-row-gap: 16px;
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
            min-width: 160px;
          }
        `}
      </style>
    </>
  );
};
