import { NextApiRequest, NextApiResponse } from 'next';

// API mock /api/current
export default function handler(req: NextApiRequest, res: NextApiResponse): void {
  res.statusCode = 200;
  res.setHeader('Content-Type', 'application/json');
  res.end(JSON.stringify({
    'temperature': getRandomTemperature(),
    'humidity': getRandomHumidity(),
    'timestamp': new Date().toISOString()
  }));
}

export const getRandomTemperature = ():number => Math.ceil(Math.random() * 200) / 10;
export const getRandomHumidity = ():number => Math.ceil(Math.random() * 80);
