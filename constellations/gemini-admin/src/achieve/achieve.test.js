import React from "react";
import Enzyme, { shallow } from "enzyme";
import { AchievePage } from "./achieve.js";

describe("test", () => {
    const component = shallow(<AchievePage />);

    test("renders", () => {
        expect(component.exists()).toBe(true);
        expect(component.find("h1").text()).toContain("All Achievements");
        expect(component.find("Link").text()).toContain("Add Achievement");
        expect(component.find("AchieveRow").length).toBe(0);
    });

    test("renders 2 rows", () => {
        const achievements = [
            {
                Id: 1,
                year: 2020,
                message: "Awesome",
            },
            {
                Id: 2,
                year: 2019,
                message: "Possum",
            },
        ];
        component.setState({ list: achievements });
        expect(component.find("AchieveRow").length).toBe(2);

        let row0 = component.find("AchieveRow").at(0);
        expect(row0.prop("achieve")).toHaveProperty("Id", 1);
        expect(row0.prop("achieve")).toHaveProperty("year", 2020);
        expect(row0.prop("achieve")).toHaveProperty("message", "Awesome");

        let row1 = component.find("AchieveRow").at(1);
        expect(row1.prop("achieve")).toHaveProperty("Id", 2);
        expect(row1.prop("achieve")).toHaveProperty("year", 2019);
        expect(row1.prop("achieve")).toHaveProperty("message", "Possum");
    });
});
