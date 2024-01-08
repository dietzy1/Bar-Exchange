/** @format */

import React from "react";
import { View } from "react-native";
import { Text } from "react-native";
import Header from "../components/Header";
import ActiveEvent from "../components/ActiveEvent";
import Timer from "../components/Timer";

const Home = () => {
  return (
    <>
      <View>
        <Header />
        <ActiveEvent />

        <View className="flex flex-row justify-center align-middle">
          <Timer />
        </View>

        <Text></Text>
      </View>
    </>
  );
};

export default Home;
