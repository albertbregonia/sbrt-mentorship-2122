class Rectangle {
    constructor(length, width) {
        this.length = length;
        this.width = width;
    }
    get length() {
        return this.length;
    }

    get width() {
        return this.width;
    }

    set length(length) {
        this.length = length; 
    }

    set width(width) {
        this.width = width;
    }
}

const rect = new Rectangle(5,7);