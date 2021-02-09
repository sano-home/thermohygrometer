import { NextApiRequest, NextApiResponse } from 'next';

import { getRandomTemperature, getRandomHumidity } from './current';

// API mock /api/histories
export default function handler(req: NextApiRequest, res: NextApiResponse): void {
  res.statusCode = 200;
  res.setHeader('Content-Type', 'application/json');

  const now: number = new Date().getTime();

  const arr = Array.of(1,2,3,4,5,6,7,8,9,10,11,12);
  const data = arr.map((item, index) => (
    {
      'temperature': getRandomTemperature(),
      'humidity': getRandomHumidity(),
      'timestamp': new Date(now - 600000 * index).toISOString()
    }
  ));


  res.end(JSON.stringify(
    {
      'pages': {
        'total': 12
      },
      'data': data.reverse()
    }
  ));
}
