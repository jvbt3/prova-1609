import 'package:flutter/material.dart';
import 'package:provaesoft8s/model/model.dart';
import 'package:provaesoft8s/service/service.dart';
import 'package:provaesoft8s/view/update_view.dart';

class ViagemView extends StatefulWidget {
  const ViagemView({super.key});

  @override
  State<ViagemView> createState() => _ViagemViewState();
}

class _ViagemViewState extends State<ViagemView> {
  final ViagemService _viagemService = ViagemService();
  late Future<List<ViagemModel>> _viagensFuture;

  @override
  void initState() {
    super.initState();
    _viagensFuture = _viagemService.fetchData();
  }

  void _refreshUsers() {
    setState(() {
      _viagensFuture = _viagemService.fetchData();
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(),
      body: FutureBuilder<List<ViagemModel>>(
        future: _viagensFuture,
        builder: (context, snapshot) {
          if (snapshot.connectionState == ConnectionState.waiting) {
            return const Center(
              child: CircularProgressIndicator(),
            );
          } else if (snapshot.hasError) {
            return const Center(
              child: Text('Nenhum usuário encontrado'),
            );
          } else if (!snapshot.hasData || snapshot.data!.isEmpty) {
            return const Center(
              child: Text('Nenhum usuário encontrado'),
            );
          } else {
            List<ViagemModel> viagens = snapshot.data!;
            return ListView.builder(
              itemCount: viagens.length,
              itemBuilder: (context, index) {
                return SizedBox(
                  height: 50,
                  child: Card(
                    child: Row(
                      mainAxisAlignment: MainAxisAlignment.spaceBetween,
                      children: [
                        Text(
                          "Name: ${viagens[index].name}",
                        ),
                        Text(
                          "Name: ${viagens[index].dataSaida}",
                        ),
                        Text(
                          "Name: ${viagens[index].dataChegada}",
                        ),
                        Text(
                          "Name: ${viagens[index].valor.toStringAsFixed(2)}",
                        ),
                        Row(
                          children: [
                            IconButton(
                              onPressed: () async {
                                int id = viagens[index].id;
                                await _viagemService.deleteData(id: id);
                                _refreshUsers();
                              },
                              icon: const Icon(Icons.delete),
                            ),
                            IconButton(
                              onPressed: () async {
                                int? id = viagens[index].id;
                                await Navigator.push(
                                  context,
                                  MaterialPageRoute(
                                    builder: (context) => UpdateView(
                                      userId: id,
                                      name: viagens[index].name,
                                      dataChegada: viagens[index].dataChegada,
                                      dataSaida: viagens[index].dataSaida,
                                      valor: viagens[index].valor,
                                    ),
                                  ),
                                );
                                _refreshUsers();
                              },
                              icon: const Icon(Icons.edit),
                            ),
                          ],
                        ),
                      ],
                    ),
                  ),
                );
              },
            );
          }
        },
      ),
    );
  }
}
