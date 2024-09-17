import 'package:flutter/material.dart';
import 'package:provaesoft8s/service/service.dart';
import 'package:provaesoft8s/view/viagem_view.dart';

class HomeView extends StatefulWidget {
  const HomeView({super.key});

  @override
  State<HomeView> createState() => _HomeViewState();
}

class _HomeViewState extends State<HomeView> {
  final UserService _userService = UserService();
  var nameController = TextEditingController();
  var dataChegadaController = TextEditingController();
  var dataSaidaController = TextEditingController();
  var valorController = TextEditingController();
  var idController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Adicionar novo usuário'),
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            SizedBox(
              width: 150,
              child: TextFormField(
                controller: idController,
                decoration: const InputDecoration(labelText: 'Id'),
              ),
            ),
            SizedBox(
              width: 150,
              child: TextFormField(
                controller: nameController,
                decoration: const InputDecoration(labelText: 'Nome'),
              ),
            ),
            SizedBox(
              width: 150,
              child: TextFormField(
                controller: dataSaidaController,
                decoration: const InputDecoration(labelText: 'Data da Saída'),
              ),
            ),
            SizedBox(
              width: 150,
              child: TextFormField(
                controller: dataChegadaController,
                decoration: const InputDecoration(labelText: 'Data da Chegada'),
              ),
            ),
            SizedBox(
              width: 150,
              child: TextFormField(
                controller: valorController,
                decoration: const InputDecoration(labelText: 'Valor'),
              ),
            ),
            Padding(
              padding: const EdgeInsets.all(8.0),
              child: ElevatedButton(
                onPressed: () {
                  final valor = valorController.text;
                  _userService.postData(
                      valor: double.parse(valor),
                      dataChegada: dataChegadaController.text,
                      dataSaida: dataSaidaController.text,
                      name: nameController.text,
                      id: int.tryParse(idController.text) ?? 0);
                  nameController.clear();
                  idController.clear();
                  dataChegadaController.clear();
                  dataSaidaController.clear();
                  valorController.clear();
                },
                child: const Text('Enviar'),
              ),
            ),
            Padding(
              padding: const EdgeInsets.all(8.0),
              child: ElevatedButton(
                onPressed: () {
                  Navigator.push(
                    context,
                    MaterialPageRoute(builder: (context) => const ViagemView()),
                  );
                },
                child: const Text('Ir para listagem'),
              ),
            )
          ],
        ),
      ),
    );
  }
}
