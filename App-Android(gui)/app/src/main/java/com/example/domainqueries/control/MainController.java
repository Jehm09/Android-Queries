package com.example.domainqueries.control;

import android.util.Log;
import android.view.View;

import com.bumptech.glide.Glide;
import com.example.domainqueries.R;
import com.example.domainqueries.model.Domain;
import com.example.domainqueries.model.History;
import com.example.domainqueries.util.Constants;
import com.example.domainqueries.util.HTTPWebUtilDomi;
import com.example.domainqueries.view.HistoryAdapter;
import com.example.domainqueries.view.MainActivity;
import com.example.domainqueries.view.ServersAdapter;
import com.google.gson.Gson;

public class MainController implements HTTPWebUtilDomi.OnResponseListener, View.OnClickListener {
    private MainActivity mainActivity;
    private HTTPWebUtilDomi util;

    public MainController(MainActivity mainActivity) {
        this.mainActivity = mainActivity;
        this.util = new HTTPWebUtilDomi();
        this.util.setListener(this);

        mainActivity.getSearchBtn().setOnClickListener(this);
        addHistory();
    }

    @Override
    public void onResponse(int callbackID, String response) {
        switch (callbackID) {
            case Constants.HOST_CALLBACK:
                Gson gsonDomain = new Gson();
                Domain domain = gsonDomain.fromJson(response, Domain.class);

                Log.e(">>>>>>", domain.getTitle());
                mainActivity.runOnUiThread(
                        () ->{
                            Glide.with(mainActivity).load(domain.getLogo()).centerCrop().into(mainActivity.getLogoIV());
                            mainActivity.getTitleTV().setText(domain.getTitle());
                            mainActivity.getSslGradeTV().setText(domain.getSsl_grade());
                            mainActivity.getPreviousSslTV().setText(domain.getPrevious_grade());
                            mainActivity.getIsDownTV().setText(domain.isIs_down()+"");
                            mainActivity.getServerChangedTV().setText(domain.isServers_changed()+"");

                            if (domain.getServers() != null) {
                                ServersAdapter serverAdapter = new ServersAdapter(domain.getServers());
                                mainActivity.getServersRecycler().setAdapter(serverAdapter);
                            }
                        }
                );

                break;
            case Constants.HISTORY_CALLBACK:
                Gson gsonHistory = new Gson();
                History history = gsonHistory.fromJson(response, History.class);

                Log.e(">>>>>>>>>>>>", history.getItems().length+"");
                mainActivity.runOnUiThread(
                        () ->{
                            if (history.getItems() != null) {
                                HistoryAdapter historyAdapter = new HistoryAdapter(history.getItems());
                                mainActivity.getHistoryRecycler().setAdapter(historyAdapter);
                            }
                        }
                );

                break;


        }
    }

    @Override
    public void onClick(View v) {
        switch (v.getId()) {
            case R.id.searchBtn:
                String hostName = mainActivity.getHostNamePT().getText().toString();

                new Thread(
                        () -> {
                            util.GETrequest(Constants.HOST_CALLBACK, Constants.URL_DOMAIN + hostName);
                        }
                ).start();

                // Updated history
                addHistory();

                break;


        }
    }


    public void addHistory() {
        new Thread(
                () -> {
                    util.GETrequest(Constants.HISTORY_CALLBACK, Constants.URL_History);
                }
        ).start();
    }
}
