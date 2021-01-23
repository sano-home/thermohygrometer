import { FC } from 'react';

export const AppContainer: FC = ({ children }) => {
  return (
    <div className="container">
      {children}
      <style jsx>{`
        .container {
          margin: 32px;
          padding: 0 0.5rem;
          display: flex;
          flex-direction: column;
          justify-content: center;
          align-items: center;
        }
      `}</style>
    </div>
  );
};
