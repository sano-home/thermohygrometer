import { FC, useEffect, useRef } from 'react';

export const ChartContainer: FC = ({ children }) => {
  const containerRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    containerRef.current.scrollLeft = 800;
  }, []);

  return (
    <div className="chart-container" ref={containerRef}>
      <div className="chart-inner">{children}</div>
      <style jsx>{`
        .chart-container {
          height: 320px;
          width: 100%;
          margin: 16px 0;
          background-color: #fff;
          border-radius: 8px;
          overflow-x: scroll;
          overflow-y: hidden;
        }

        .chart-inner {
          width: 768px;
          height: 320px;
        }
      `}</style>
    </div>
  );
};
