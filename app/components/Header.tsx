/** @format */
import { View } from "react-native";
import { Text } from "react-native";

import React from "react";

const Header = () => {
  return (
    <>
      <View className="flex flex-row items-center py-12 justify-center">
        <Text className="text-green-500 text-3xl underline">BØRS</Text>
        <Text className="text-red-500 text-3xl underline">BAR</Text>
      </View>
    </>
  );
};

export default Header;

{
  /* <div className="bg-gray-900 text-white">
<div className="flex flex-col items-center py-8">
  <h1 className="text-4xl md:text-6xl lg:text-7xl font-bold">
    <AiOutlineStock className="inline-block mr-2 text-green-500" />

    <span className="text-green-500">BØRS</span>
    <span className="text-red-500">BAR</span>

    <AiOutlineStock
      className="inline-block ml-2 text-red-500"
      style={{ transform: "scaleX(-1)" }}
    />
  </h1>
  <p className="mt-2 text-lg">Udbud og efterspørgsel styrer prisen</p>
</div>
</div> */
}
