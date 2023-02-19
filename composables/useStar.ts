interface StarInterface {
    Id: number;
    Name: string;
    Avatar: string;
    CreateTime: string;
    Content: string;
}

export interface UserInterface {
    ids: string;
    stars: StarInterface[];
}

export const useStar = () => useState<UserInterface>("star", () => {
    return {
        ids: "",
        stars: []
    }
})