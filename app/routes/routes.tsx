/** @format */

import { createMaterialBottomTabNavigator } from "react-native-paper/react-navigation";
import MaterialCommunityIcons from "react-native-vector-icons/MaterialCommunityIcons";

/// <reference types="nativewind/types" />

import Home from "../views/Home";
import Event from "../views/Event";
import Drinks from "../views/Drinks";

const Tab = createMaterialBottomTabNavigator();
function Navbar() {
  //#22C55E
  return (
    <Tab.Navigator
      initialRouteName="Home"
      activeColor="#EF4444"
      barStyle={{ backgroundColor: "#FFFFFF" }}
    >
      <Tab.Screen
        name="Home"
        component={Home}
        options={{
          tabBarLabel: "Home",
          tabBarIcon: ({ color }) => (
            <MaterialCommunityIcons name="home" color={"#EF4444"} size={26} />
          ),
        }}
      />
      <Tab.Screen
        name="Event"
        component={Event}
        options={{
          tabBarLabel: "Event",
          tabBarIcon: () => (
            <MaterialCommunityIcons name="bell" color={"#EF4444"} size={26} />
          ),
        }}
      />
      <Tab.Screen
        name="Drinks"
        component={Drinks}
        options={{
          tabBarLabel: "Drinks",
          tabBarIcon: () => (
            <MaterialCommunityIcons name="beer" color={"#EF4444"} size={26} />
          ),
        }}
      />
    </Tab.Navigator>
  );
}

export default Navbar;
