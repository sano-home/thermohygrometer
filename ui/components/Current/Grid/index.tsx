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
            width: 800px;
          }

          @media (max-width: 767px) {
            .grid {
              width: 100%;
              flex-direction: column;
            }
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
            width: 384px;
          }

          @media (max-width: 767px) {
            .grid-item {
              margin-bottom: 16px;
            }
          }
        `}
      </style>
    </>
  );
};
