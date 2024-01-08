/** @format */

import React from "react";
import { View, Text } from "react-native";
import { Button } from "react-native-paper";
import { TimePickerModal } from "react-native-paper-dates";
import { SafeAreaProvider } from "react-native-safe-area-context";

const TimePicker = () => {
  const [hours, setHours] = React.useState(0); // [1
  const [minutes, setMinutes] = React.useState(0);

  const [visible, setVisible] = React.useState(false);
  const onDismiss = React.useCallback(() => {
    setVisible(false);
  }, [setVisible]);

  const onConfirm = React.useCallback(
    ({ hours, minutes }: { hours: number; minutes: number }) => {
      setVisible(false);
      setHours(hours);
      setMinutes(minutes);

      console.log({ hours, minutes });
    },
    [setVisible]
  );

  return (
    <SafeAreaProvider>
      <View style={{ justifyContent: "center", flex: 1, alignItems: "center" }}>
        <Text>
          {hours}
          {minutes}
        </Text>
        <Button
          onPress={() => setVisible(true)}
          uppercase={false}
          mode="outlined"
        >
          Pick time
        </Button>
        <TimePickerModal
          visible={visible}
          onDismiss={onDismiss}
          onConfirm={onConfirm}
          hours={12}
          minutes={14}
        />
      </View>
    </SafeAreaProvider>
  );
};

export default TimePicker;

//Either we display the current event or we display
