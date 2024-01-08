/** @format */

import React from "react";
import { View, Text } from "react-native";
import { Button } from "react-native-paper";
import { TimePickerModal } from "react-native-paper-dates";
import { SafeAreaProvider } from "react-native-safe-area-context";
import DatePicker from "../components/DatePicker";
import TimePicker from "../components/TimePicker";

const Event = () => {
  return (
    <>
      <DatePicker />
      <TimePicker />
    </>
  );
};

export default Event;

//Either we display the current event or we display
