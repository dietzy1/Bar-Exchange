/** @format */

import React from "react";
import { useState, useEffect } from "react";
import { View, ScrollView, Text } from "react-native";
import { Button, Card, Title, Paragraph } from "react-native-paper";
import useGetBeverages from "../api/endpoints/beverage/getBeverages";
import Header from "../components/Header";
import BeerSvg from "../components/BeerSvg";
import CocktailSvg from "../components/CocktailSvg";
import ShotSvg from "../components/ShotSvg";
import { ActivityIndicator } from "react-native";
import CardCover from "react-native-paper/lib/typescript/components/Card/CardCover";

const Drinks = () => {
  const { data, isLoading, error } = useGetBeverages();
  const [selectedCategory, setSelectedCategory] = useState("beer"); // Default category is "beer"

  // Filter data based on the selected category
  const filteredData = data?.beverages.filter((beverage) => {
    return (
      beverage.type.toString() ===
      `BEVERAGE_TYPE_${selectedCategory.toUpperCase()}`
    );
  });

  return (
    <View>
      <Header />
      <View className="flex flex-row justify-around px-4 bg-red-500 py-2">
        {/* Category buttons */}

        <Button
          mode="contained"
          icon={BeerSvg}
          onPress={() => setSelectedCategory("beer")}
          className={` ${
            selectedCategory === "beer" ? "bg-green-500" : "bg-gray-400"
          }`}
        >
          Øl
        </Button>
        <Button
          mode="contained"
          icon={CocktailSvg}
          onPress={() => setSelectedCategory("cocktail")}
          className={`${
            selectedCategory === "cocktail" ? "bg-green-500" : "bg-gray-400"
          }`}
        >
          Cocktails
        </Button>
        <Button
          mode="contained"
          icon={ShotSvg}
          onPress={() => setSelectedCategory("shots")}
          className={`${
            selectedCategory === "shots" ? "bg-green-500" : "bg-gray-400"
          }`}
        >
          Shots
        </Button>
      </View>

      <ScrollView>
        <View className="px-8 py-4">
          {isLoading ? (
            <ActivityIndicator size="large" color="#00ff00" />
          ) : error ? (
            <ActivityIndicator size="large" color="#00ff00" />
          ) : (
            filteredData?.map((item) => (
              <Card key={item.id} className="mb-4">
                <Card.Content className="flex flex-row items-center py-6">
                  <BeerSvg />
                  <Title className="ml-2">{item.name}</Title>
                  <View style={{ flex: 1 }} />
                  {/* This will create a flexible space */}
                  <Paragraph className="mr-2">{item.price} dkk</Paragraph>
                  <Button className="ml-auto" textColor="red" mode="outlined">
                    Edit
                  </Button>
                </Card.Content>
              </Card>
            ))
          )}
        </View>
      </ScrollView>
      <View className="bg-green-500 flex flex-row justify-center my-auto py-4 absolute bottom-0 w-full">
        <Button
          mode="contained"
          icon={BeerSvg}
          onPress={() => setSelectedCategory("beer")}
          className={` ${
            selectedCategory === "beer" ? "bg-green-500" : "bg-gray-400"
          }`}
        >
          Tilføj
        </Button>
      </View>
    </View>
  );
};

export default Drinks;
