angular.module("items", [])
.service("Items", function($http) {
    this.getRecent = function(count) {
        return new Promise(function(resolve, reject) {
            resolve([cat, cat, cat])
        });
    };

    return this;
});

const cat = {
    id: "abc123",
    name: "This Cat",
    description: "This smiling cat that is smiling.",
    categories: [
        {
            id: "cat-animals",
            name: "Animals",
        },
        {
            id: "cat-cats",
            name: "Cats"
        }
    ],
    location: {
        id: "loc-living-room",
        name: "Living Room"
    },
    quantity: 1,
    image: {
        full: "https://s3.us-east-2.amazonaws.com/alex-stuff-tracker/full/smilecat.png",
        thumb: "https://s3.us-east-2.amazonaws.com/alex-stuff-tracker/full/smilecat.png"
    }
};
