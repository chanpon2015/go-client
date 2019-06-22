class Menu {
    private id: string;
    private icon: string;
    private name: string;
    constructor(id: string, icon: string, name: string) {
        this.id = id;
        this.icon = icon;
        this.name = name;
    }
    public getId(): string {
        return this.id
    }
    public getIcon(): string {
        return this.icon;
    }
    public getName(): string {
        return this.name;
    }
}