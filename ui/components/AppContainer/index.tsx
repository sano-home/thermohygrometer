import { FC } from 'react';

export const AppContainer: FC = ({ children }) => {
  return (
    <div className="container">
      {children}
      <style jsx>{`
        .container {
          margin: auto;
          padding: 1rem;
          display: flex;
          flex-direction: column;
          justify-content: center;
          align-items: center;
          max-width: 800px;
        }
      `}</style>
    </div>
  );
};
