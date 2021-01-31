import { HistoryChart } from '../components/HistoryChart';
import { Current } from '../components/Current/index';

export default function Home(): JSX.Element {
  return (
    <>
      <Current />
      <HistoryChart />
    </>
  );
}
