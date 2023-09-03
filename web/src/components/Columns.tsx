/** @format */

"use client";

import { ColumnDef } from "@tanstack/react-table";
import { ArrowUpIcon, ArrowDownIcon } from "@heroicons/react/24/solid";
import { Beverage } from "@/api/protos/beverage/v1/beverage_pb";

export const columns: ColumnDef<Beverage>[] = [
  {
    accessorKey: "name",
    header: () => (
      <div className="">
        <div className="font-medium text-white">Type</div>
      </div>
    ),
  },

  {
    accessorKey: "price",
    header: () => <div className="text-right text-white">Pris</div>,
    cell: ({ row }) => {
      const price = parseFloat(row.getValue("price"));

      const formatted = new Intl.NumberFormat("en-US", {
        style: "currency",
        currency: "dkk",
      }).format(price);

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

      if (status === "STATUS_INCREASING") {
        return (
          <div className="flex flex-row">
            <ArrowDownIcon className="text-red-500 h-6" />
          </div>
        );
      } else if (status === "STATUS_DECREASING") {
        return (
          <div className="flex flex-row">
            <ArrowUpIcon className="text-green-500 h-6" />
          </div>
        );
      } else if (status === "STATUS_NO_CHANGE")
        return (
          <div className="flex flex-row">
            {/*  <ArrowDownIcon className="text-black h-6" /> */}
          </div>
        );

      {
        return <div></div>;
      }
    },
  },

  {
    accessorKey: "percentageChange",
    header: () => <div className="text-right text-white">%</div>,
    cell: ({ row }) => {
      const PercentageChange = parseFloat(row.getValue("percentageChange"));

      //check if percentage is negative
      if (PercentageChange < 0) {
        return (
          <div className="text-right font-medium flex flex-row justify-end text-green-500">
            <div className="">{PercentageChange.toFixed(2)}%</div>
          </div>
        );
      }

      if (PercentageChange === 0) {
        return (
          <div className="text-right font-medium flex flex-row justify-end text-black">
            <div className="">{PercentageChange.toFixed(2)}%</div>
          </div>
        );
      }

      return (
        <div className="text-right font-medium flex flex-row justify-end text-red-500">
          <div className="">{PercentageChange.toFixed(2)}%</div>
        </div>
      );
    },
  },
  // Add other columns as needed
];
