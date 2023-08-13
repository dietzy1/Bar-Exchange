/** @format */

import { useEffect, useState } from "react";

export default function Timer({ totalHours }: { totalHours: number }) {
  const [remainingTime, setRemainingTime] = useState(totalHours * 60 * 60);

  useEffect(() => {
    const interval = setInterval(() => {
      setRemainingTime((prevTime) => {
        const newTime = prevTime - 1;
        return newTime < 0 ? 0 : newTime;
      });
    }, 1000);

    return () => clearInterval(interval);
  }, [totalHours]);

  const formatTime = (seconds: number) => {
    const hours = Math.floor(seconds / 3600);
    const minutes = Math.floor((seconds % 3600) / 60);
    const remainingSeconds = seconds % 60;
    return {
      hours: hours.toString().padStart(2, "0"),
      minutes: minutes.toString().padStart(2, "0"),
      seconds: remainingSeconds.toString().padStart(2, "0"),
    };
  };

  const { hours, minutes, seconds } = formatTime(remainingTime);

  return (
    <div className="flex justify-center space-x-4">
      <div className="flex flex-col items-center w-16 h-16  bg-white">
        <div className="text-2xl font-bold text-black">{hours}</div>
        <div className="text-gray-600 text-sm text-center">Timer</div>
      </div>
      <div className="flex flex-col items-center w-16 h-16  bg-white">
        <div className="text-2xl font-bold text-black">{minutes}</div>
        <div className="text-gray-600 text-sm text-center">Minutter</div>
      </div>
      <div className="flex flex-col items-center w-16 h-16  bg-white">
        <div className="text-2xl font-bold text-red-500">{seconds}</div>
        <div className="text-gray-600 text-sm text-center">Sekunder</div>
      </div>
    </div>
  );
}
