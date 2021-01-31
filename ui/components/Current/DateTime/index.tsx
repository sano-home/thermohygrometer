import { FC } from 'react';

import { GridContainer, GridItem } from '../Grid';

export const DateTime: FC<{ timestamp: string }> = ({ timestamp }) => {
  const currentTime = new Date(timestamp);

  return (
    <div className="container">
      <GridContainer>
        <GridItem>
          <h1>{currentTime.toLocaleDateString()}</h1>
        </GridItem>
        <GridItem>
          <h1 style={{ textAlign: 'right' }}>
            {currentTime.toLocaleTimeString().substring(0, 5)}
          </h1>
        </GridItem>
      </GridContainer>

      <style jsx>{`
        .container {
          width: 100%;
        }

        h1 {
          margin: 0;
        }
      `}</style>
    </div>
  );
};
