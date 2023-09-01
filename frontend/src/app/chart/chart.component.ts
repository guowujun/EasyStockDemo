import { Component, ElementRef, OnInit } from '@angular/core';
import { createChart } from "lightweight-charts";
import { GetKline, GetupDown } from "../../../wailsjs/go/main/App";

import { faArrowTrendUp, faArrowTrendDown } from "@fortawesome/free-solid-svg-icons";

@Component({
    selector: 'app-chart',
    templateUrl: './chart.component.html',
    styleUrls: ['./chart.component.css']
})
export class ChartComponent implements OnInit {
    trendUpIcon = faArrowTrendUp;
    TrendDownIcon = faArrowTrendDown;

    chart: any;
    candlestickSeries: any;
    chartOptions: any;
    bgColor = "#2B2D30";
    footerData: any;
	disabled = false;

    constructor(private el: ElementRef) {
    }

    ngOnInit(): void {
        this.chartOptions = {
            width: 800,
            height: 600,
            layout: {
                background: {color: this.bgColor},
                textColor: "#C3BCDB",
            },
            grid: {
                vertLines: {color: "#444", visible: false,},
                horzLines: {color: "#444"},
            },
            timeScale: {
                borderColor: '#71649C',
            },
            localization: {
                dateFormat: "yyyy-MM-dd",
            }
        };
        this.chart = createChart(this.el.nativeElement.querySelector("#container"), this.chartOptions);
        this.chart.timeScale().fitContent();
        this.candlestickSeries = this.chart.addCandlestickSeries({
            upColor: '#ef5350',
            wickUpColor: '#ef5350',
            downColor: '#26a69a',
            wickDownColor: '#26a69a',
            borderVisible: false,
            priceLineVisible: false,
        });

        this.candlestickSeries.priceScale().applyOptions({
            borderColor: '#71649C',
            scaleMargins: {
                top: 0.1,
                bottom: 0.3,
            },
        });
    }

    renderChart(data: any) {
        this.candlestickSeries.setData(
            data
        );
    }

    do() {
		this.disabled = true;
        setInterval(() => {
            GetKline("600822", 200).then(value => {
                this.renderChart(value);
                console.log(value);
            });

            GetupDown().then(value => {
                console.log(value);
                this.footerData = value;
            })

        }, 5000);
    }

}
