import { FC } from 'react';

export const ChartContainer: FC = ({ children }) => {
  return (
    <div className="chart-container">
      {children}
      <style jsx>{`
        .chart-container {
          height: 320px;
          width: 100%;
          margin: 16px 0;
          background-color: #fff;
          border-radius: 8px;
        }
      `}</style>
    </div>
  );
};
