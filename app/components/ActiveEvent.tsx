/** @format */

import React from "react";
import { View } from "react-native";
import { Text } from "react-native";
import ENV from "../constants/env";
import { Button, Card, Title, Paragraph } from "react-native-paper";

const ActiveEvent = () => {
  //Figure out if the event is active or not

  const isActive = true; // Replace with your logic to determine if the event is active

  return (
    <View className="mx-8">
      <Card className="bg-white rounded-lg shadow-lg">
        <Card.Content>
          <Title className="text-xl font-bold">Active Event</Title>
          {isActive ? (
            <Paragraph className="text-green-500">
              There is an active event right now.
            </Paragraph>
          ) : (
            <Paragraph className="text-red-500">
              There are no active events at the moment.
            </Paragraph>
          )}
        </Card.Content>

        {isActive && (
          <Card.Actions>
            <Button
              className="bg-green-500 text-white "
              mode="contained"
              onPress={() => {
                // Handle navigation to the event details screen
              }}
            >
              View Event
            </Button>
          </Card.Actions>
        )}
      </Card>
    </View>
  );
};

export default ActiveEvent;
