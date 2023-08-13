/** @format */

"use client";

import { ColumnDef } from "@tanstack/react-table";
import { ArrowUpIcon, ArrowDownIcon } from "@heroicons/react/24/solid";

export type Beverage = {
  id: string;
  item: string;
  kind: string;
  amount: number;
  PercentageChange: number;
  status: "increasing" | "decreasing" | "no_change";
};

export const columns: ColumnDef<Beverage>[] = [
  {
    accessorKey: "item",
    header: () => (
      <div className="">
        <div className="font-medium text-white">Type</div>
      </div>
    ),
  },

  {
    accessorKey: "amount",
    header: () => <div className="text-right text-white">Pris</div>,
    cell: ({ row }) => {
      const amount = parseFloat(row.getValue("amount"));

      const formatted = new Intl.NumberFormat("en-US", {
        style: "currency",
        currency: "dkk",
      }).format(amount);

      return (
        <div className="text-right font-medium flex flex-row justify-end">
          <div className="">{formatted}</div>
        </div>
      );
    },
  },

  {
    accessorKey: "status",
    header: () => <div></div>,
    cell: ({ row }) => {
      const status = row.getValue("status");

      if (status === "increasing") {
        return (
          <div className="flex flex-row">
            <ArrowUpIcon className="text-green-500 h-6" />
          </div>
        );
      } else if (status === "decreasing") {
        return (
          <div className="flex flex-row">
            <ArrowDownIcon className="text-red-500 h-6" />
          </div>
        );
      } else {
        return <div></div>;
      }
    },
  },

  {
    accessorKey: "PercentageChange",
    header: () => <div className="text-right text-white">1h%</div>,
    cell: ({ row }) => {
      const PercentageChange = parseFloat(row.getValue("PercentageChange"));

      const formatted = new Intl.NumberFormat("en-US", {
        style: "percent",
        maximumFractionDigits: 2,
      }).format(PercentageChange);

      return (
        <div className="text-right font-medium flex flex-row justify-end text-red-500">
          <div className="">{formatted}</div>
        </div>
      );
    },
  },
  // Add other columns as needed
];

//Categories I want to display
//- beer
//- drinks
//- Shots

//Data I want to display
//- name
//- price
//- up/down arrow
//- hourly persentage change

//I probaly want to display 2 tables
