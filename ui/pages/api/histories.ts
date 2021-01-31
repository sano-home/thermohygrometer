import { NextApiRequest, NextApiResponse } from 'next';

export default function handler(req: NextApiRequest, res: NextApiResponse): void {
  res.statusCode = 200;
  res.setHeader('Content-Type', 'application/json');
  res.end(JSON.stringify(
    {
      'pages': {
        'total': 3
      },
      'data': [
        {
          'temperature': 22,
          'humidity': 12,
          'timestamp': '2021-01-02T11:37:42.337Z'
        },
        {
          'temperature': 23,
          'humidity': 13,
          'timestamp': '2021-01-02T11:38:42.337Z'
        },
        {
          'temperature': 24,
          'humidity': 14,
          'timestamp': '2021-01-02T11:39:42.337Z'
        }
      ]
    }
  ));
}
