/** @format */
/// <reference types="nativewind/types" />

import { NavigationContainer } from "@react-navigation/native";
import { SafeAreaProvider } from "react-native-safe-area-context";

import { createNativeStackNavigator } from "@react-navigation/native-stack";

import { StatusBar } from "expo-status-bar";
import { StyleSheet, Text, View } from "react-native";
import Navbar from "./routes/routes";
import ReactQueryClientProvider from "./api/queryClient";

export default function App() {
  return (
    <SafeAreaProvider>
      <ReactQueryClientProvider>
        <NavigationContainer>
          <Navbar />
        </NavigationContainer>
      </ReactQueryClientProvider>
    </SafeAreaProvider>
  );
}

/* const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: "#fff",
    alignItems: "center",
    justifyContent: "center",
  },
});s
 */
