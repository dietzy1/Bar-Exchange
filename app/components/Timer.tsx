/** @format */

import React, { useEffect, useState } from "react";
import { View, Text } from "react-native";
import useGetEvent from "../api/endpoints/event/getEvent";

function calculateTimeDifference(targetDateTime: Date, currentDateTime: Date) {
  const timeDifferenceMilliseconds =
    targetDateTime.getTime() - currentDateTime.getTime();
  const totalSecondsRemaining = Math.floor(timeDifferenceMilliseconds / 1000);
  return totalSecondsRemaining;
}

export default function Timer(): JSX.Element {
  const { data, isLoading, error } = useGetEvent();

  const [remainingTime, setRemainingTime] = useState(0);

  useEffect(() => {
    if (!isLoading && data) {
      const targetDateTime = new Date(data.futureTimestamp);
      const currentDateTime = new Date();

      const initialRemainingTime = calculateTimeDifference(
        targetDateTime,
        currentDateTime
      );
      setRemainingTime(initialRemainingTime);

      const interval = setInterval(() => {
        setRemainingTime((prevTime) => {
          const newTime = prevTime - 1;
          return newTime < 0 ? 0 : newTime;
        });
      }, 1000);

      return () => clearInterval(interval);
    }
  }, [data, isLoading]);

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

  if (error) {
    console.log(error);
    return (
      <View className="flex justify-center space-x-4 invisible">
        <View className="flex flex-col items-center w-16 h-16  bg-white">
          <div className="text-2xl font-bold text-black">{}</div>
          <div className="text-gray-600 text-sm text-center">Timer</div>
        </View>
        <View className="flex flex-col items-center w-16 h-16  bg-white">
          <div className="text-2xl font-bold text-black">{}</div>
          <div className="text-gray-600 text-sm text-center">Minutter</div>
        </View>
        <View className="flex flex-col items-center w-16 h-16  bg-white">
          <div className="text-2xl font-bold text-red-500">{}</div>
          <div className="text-gray-600 text-sm text-center">Sekunder</div>
        </View>
      </View>
    );
  }

  //Main component render
  return (
    <View className="flex justify-center flex-row space-x-4">
      <View className="flex flex-col items-center w-16 h-16  bg-white">
        <Text className="text-2xl font-bold text-black">{hours}</Text>
        <Text className="text-gray-600 text-sm text-center">Timer</Text>
      </View>
      <View className="flex flex-col items-center w-16 h-16  bg-white">
        <Text className="text-2xl font-bold text-black">{minutes}</Text>
        <Text className="text-gray-600 text-sm text-center">Minutter</Text>
      </View>
      <View className="flex flex-col items-center w-16 h-16  bg-white">
        <Text className="text-2xl font-bold text-red-500">{seconds}</Text>
        <Text className="text-gray-600 text-sm text-center">Sekunder</Text>
      </View>
    </View>
  );
}
