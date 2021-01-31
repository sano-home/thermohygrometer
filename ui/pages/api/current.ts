import { NextApiRequest, NextApiResponse } from 'next';

export default function handler(req: NextApiRequest, res: NextApiResponse): void {
  res.statusCode = 200;
  res.setHeader('Content-Type', 'application/json');
  res.end(JSON.stringify({
    'temperature':22,
    'humidity':15,
    'timestamp':'2021-01-10T13:25:48Z'
  }));
}
