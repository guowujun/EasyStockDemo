export namespace main {
	
	export class IndexData {
	    name: string;
	    date: string;
	    close: number;
	    upCount: number;
	    downCount: number;
	
	    static createFrom(source: any = {}) {
	        return new IndexData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.date = source["date"];
	        this.close = source["close"];
	        this.upCount = source["upCount"];
	        this.downCount = source["downCount"];
	    }
	}
	export class Kline {
	    time: string;
	    open: number;
	    high: number;
	    low: number;
	    close: number;
	    vol: number;
	    amount: number;
	
	    static createFrom(source: any = {}) {
	        return new Kline(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.time = source["time"];
	        this.open = source["open"];
	        this.high = source["high"];
	        this.low = source["low"];
	        this.close = source["close"];
	        this.vol = source["vol"];
	        this.amount = source["amount"];
	    }
	}

}

